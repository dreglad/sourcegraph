import { psub } from '$lib/utils'
import type { PageLoad } from './$types'
import { queryGitBranchesOverview } from '@sourcegraph/web/src/repo/branches/loader'
import { isErrorLike } from '@sourcegraph/common/src/errors/utils'

export const load: PageLoad = ({ parent }) => {
    return {
        prefetch: {
            branches: psub(
                parent().then(({ resolvedRevision }) =>
                    isErrorLike(resolvedRevision)
                        ? null
                        : queryGitBranchesOverview({ repo: resolvedRevision.repo.id, first: 10 }).toPromise()
                )
            ),
        },
    }
}
