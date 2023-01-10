<script lang="ts">
    import FileTree from '$lib/FileTree.svelte'

    import Commit from '$lib/Commit.svelte'
    import {
        mdiSourceCommit,
        mdiSourceRepository,
        mdiFolderOutline,
        mdiFileDocumentOutline,
        mdiSourceBranch,
        mdiTag,
    } from '@mdi/js'
    import { isErrorLike } from '@sourcegraph/common'

    import type { PageData } from './$types'
    import Icon from '$lib/Icon.svelte'

    export let data: PageData

    $: treeOrError = data.preload.treeEntries
    $: commits = data.preload.commits
</script>

<section>
    <div class="sidebar">
        {#if !$treeOrError.loading && $treeOrError.data}
            <FileTree activeEntry="" treeOrError={$treeOrError.data} />
        {/if}
    </div>
    <div class="content">
        <h1><Icon svgPath={mdiSourceRepository} /> {data.repoName}</h1>
        {#if !isErrorLike(data.resolvedRevision)}
            <p>
                {data.resolvedRevision.repo.description}
            </p>
        {/if}
        <p class="tabs">
            <a href="{data.repoURL}/-/commits"><Icon svgPath={mdiSourceCommit} inline /> Commits</a>
            <a href="{data.repoURL}/-/branches"><Icon svgPath={mdiSourceBranch} inline /> Branches</a>
            <a href="{data.repoURL}/-/tags"><Icon svgPath={mdiTag} inline /> Tags</a>
        </p>

        {#if !$treeOrError.loading && $treeOrError.data && !isErrorLike($treeOrError.data)}
            <h3>Files and directories</h3>
            <ul class="files">
                {#each $treeOrError.data.entries as entry}
                    <li>
                        <a
                            data-sveltekit-preload-data={entry.isDirectory ? 'hover' : 'tap'}
                            data-sveltekit-preload-code="hover"
                            href={entry.url}
                            ><Icon svgPath={entry.isDirectory ? mdiFolderOutline : mdiFileDocumentOutline} inline />
                            {entry.name}</a
                        >
                    </li>
                {/each}
            </ul>
        {/if}

        <h3 class="mt-3">Changes</h3>
        <ul class="commits">
            {#if $commits.loading}
                Loading...
            {:else if $commits.data}
                {#each $commits.data as commit (commit.url)}
                    <li><Commit {commit} /></li>
                {/each}
            {/if}
        </ul>
    </div>
</section>

<style lang="scss">
    section {
        display: flex;
        overflow: hidden;
        margin: 1rem;
        margin-bottom: 0;
        flex: 1;
    }

    .sidebar {
        width: 200px;
        overflow: hidden;
        display: flex;
        flex-direction: column;
    }

    .content {
        padding: 1rem;
        flex: 1;
        overflow: auto;
        margin: 1rem;
        margin-bottom: 0;
        background-color: var(--color-bg-1);
        border-radius: var(--border-radius);
        overflow: hidden;
        display: flex;
        flex-direction: column;
        border: 1px solid var(--border-color);
        border-radius: var(--border-radius);
    }

    .tabs > a {
        padding: 0.25rem 0.5rem;
        color: var(--body-color);
        border: 1px solid var(--border-color);
        border-radius: var(--border-radius);
    }

    ul.commits {
        padding: 0;
        margin: 0;
        list-style: none;

        li {
            border-bottom: 1px solid var(--border-color);
            padding: 0.5rem 0;

            &:last-child {
                border: none;
            }
        }
    }

    ul.files {
        padding: 0;
        margin: 0;
        list-style: none;
        columns: 3;
    }
</style>
