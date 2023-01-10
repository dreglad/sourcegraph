import type { LayoutLoad } from './$types'
import { resolveRepoRevision } from '@sourcegraph/web/src/repo/backend'
import { parseRepoRevision } from '@sourcegraph/shared/src/util/url'
import { catchError } from 'rxjs/operators/index'
import { isCloneInProgressErrorLike, isRepoSeeOtherErrorLike } from '@sourcegraph/shared/src/backend/errors'
import { NEVER, of } from 'rxjs'
import { asError, encodeURIPathComponent, type ErrorLike } from '@sourcegraph/common'

export const load: LayoutLoad = ({ params }) => {
    const { repoName, revision } = parseRepoRevision(params.repo)
    const resolvedRevision = resolveRepoRevision({ repoName, revision })
        .pipe(
            catchError(error => {
                const redirect = isRepoSeeOtherErrorLike(error)

                if (redirect) {
                    //redirectToExternalHost(redirect)
                    return NEVER
                }

                if (isCloneInProgressErrorLike(error)) {
                    return of<ErrorLike>(asError(error))
                }

                throw error
            })
        )
        .toPromise()

    return {
        repoURL: '/' + encodeURIPathComponent(repoName),
        repoName,
        revision,
        resolvedRevision,
    }
}
