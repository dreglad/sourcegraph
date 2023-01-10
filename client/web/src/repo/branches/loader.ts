import { memoizeObservable } from '@sourcegraph/common/src/util/rxjs/memoizeObservable'
import type { Scalars } from '@sourcegraph/shared/src/graphql-operations'
import type { Observable } from 'rxjs'

import { gql } from '@sourcegraph/http-client'

import { queryGraphQL } from '../../backend/graphql'
import type { GitRefFields, RepositoryGitBranchesOverviewRepository } from '../../graphql-operations'
import { map } from 'rxjs/operators/index'
import { createAggregateError } from '@sourcegraph/common/src/errors/errors'
import { gitReferenceFragments } from '../loader'

interface Data {
    defaultBranch: GitRefFields | null
    activeBranches: GitRefFields[]
    hasMoreActiveBranches: boolean
}

export const queryGitBranchesOverview = memoizeObservable(
    (args: { repo: Scalars['ID']; first: number }): Observable<Data> =>
        queryGraphQL(
            gql`
                query RepositoryGitBranchesOverview($repo: ID!, $first: Int!, $withBehindAhead: Boolean!) {
                    node(id: $repo) {
                        ...RepositoryGitBranchesOverviewRepository
                    }
                }

                fragment RepositoryGitBranchesOverviewRepository on Repository {
                    defaultBranch {
                        ...GitRefFields
                    }
                    gitRefs(first: $first, type: GIT_BRANCH, orderBy: AUTHORED_OR_COMMITTED_AT) {
                        nodes {
                            ...GitRefFields
                        }
                        pageInfo {
                            hasNextPage
                        }
                    }
                }
                ${gitReferenceFragments}
            `,
            { ...args, withBehindAhead: true }
        ).pipe(
            map(({ data, errors }) => {
                if (!data || !data.node) {
                    throw createAggregateError(errors)
                }
                const repo = data.node as RepositoryGitBranchesOverviewRepository
                if (!repo.gitRefs || !repo.gitRefs.nodes) {
                    throw createAggregateError(errors)
                }
                return {
                    defaultBranch: repo.defaultBranch,
                    activeBranches: repo.gitRefs.nodes.filter(
                        // Filter out default branch from activeBranches.
                        ({ id }) => !repo.defaultBranch || repo.defaultBranch.id !== id
                    ),
                    hasMoreActiveBranches: repo.gitRefs.pageInfo.hasNextPage,
                }
            })
        ),
    args => `${args.repo}:${args.first}`
)
