import { psub } from '$lib/utils'
import { isErrorLike } from '@sourcegraph/common/src/errors/utils'
import { GitRefType } from '@sourcegraph/search'
import { queryGitReferences } from '@sourcegraph/web/src/repo/loader'
import type { PageLoad } from './$types'

export const load: PageLoad = ({ parent }) => {
    return {
        preload: {
            tags: psub(
                parent().then(({ resolvedRevision }) =>
                    isErrorLike(resolvedRevision)
                        ? null
                        : queryGitReferences({
                              repo: resolvedRevision.repo.id,
                              type: GitRefType.GIT_TAG,
                              first: 20,
                          }).toPromise()
                )
            ),
        },
    }
}
