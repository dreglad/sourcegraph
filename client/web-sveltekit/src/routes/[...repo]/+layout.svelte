<script lang="ts">
    import { page } from '$app/stores'
    import Icon from '$lib/Icon.svelte'
    import { mdiChevronRight, mdiSourceRepository } from '@mdi/js'
    import { isErrorLike } from '@sourcegraph/common'
    import type { LayoutData } from './$types'
    import { createActionContext } from '$lib/repo/actions'
    import { setContext } from 'svelte'
    import Permalink from './Permalink.svelte'
    import { getRevisionLabel, navFromPath } from '$lib/repo/utils'
    import { displayRepoName } from '$lib/shared/repo'

    export let data: LayoutData

    // Sets up a context for other components to add add buttons to the header
    const repoActions = createActionContext()
    setContext('repo-actions', repoActions)

    $: ({ repo, path } = $page.params)
    $: nav = path ? navFromPath(path, repo, $page.url.pathname.includes('/-/blob/')) : []

    $: resolvedRevision = isErrorLike(data.resolvedRevision) ? null : data.resolvedRevision
    $: revisionLabel = getRevisionLabel(data.revision, resolvedRevision)
    $: repoName = displayRepoName(repo.split('@')[0])
    $: if (resolvedRevision) {
        repoActions.setAction({ key: 'permalink', priority: 100, component: Permalink })
    }
</script>

{#if isErrorLike(data.resolvedRevision)}
    Something went wrong
{:else}
    <div class="header">
        <nav>
            <a class="button" href="/{repo}"><Icon svgPath={mdiSourceRepository} inline /> {repoName}</a>
            {#if revisionLabel}
                @ <span class="button">{revisionLabel}</span>
            {/if}
            {#if nav.length > 0}
                <Icon svgPath={mdiChevronRight} inline />
                <span class="crumps">
                    {#each nav as [label, url]}
                        <span>/</span>
                        <a href={url}>{label}</a>&nbsp;
                    {/each}
                </span>
            {/if}
        </nav>
        <div class="actions">
            {#each $repoActions as action (action.key)}
                <svelte:component this={action.component} />
            {/each}
        </div>
    </div>
    <slot />
{/if}

<style>
    .header {
        display: flex;
        align-items: center;
        margin: 1rem 1rem 0 1rem;
    }

    .actions {
        margin-left: auto;
    }

    nav {
        color: var(--body-color);
    }

    .crumps {
        color: var(--link-color);
    }

    .button {
        color: var(--body-color);
        border: 1px solid var(--border-color);
        padding: 0.25rem 0.5rem;
        border-radius: var(--border-radius);
        text-decoration: none;
    }
</style>
