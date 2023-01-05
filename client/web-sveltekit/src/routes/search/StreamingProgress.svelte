<script lang="ts">
    import Icon from '$lib/Icon.svelte'

    import Popover from '$lib/Popover.svelte'
    import SyntaxHighlightedQuery from '$lib/search/SyntaxHighlightedQuery.svelte'
    import { mdiAlertCircle, mdiChevronDown, mdiChevronLeft, mdiInformationOutline, mdiMagnify } from '@mdi/js'
    import { renderMarkdown } from '@sourcegraph/common'
    import { pluralize } from '@sourcegraph/common/src/util/strings'
    import { limitHit, sortBySeverity } from '@sourcegraph/search-ui/src/results/progress/utils'
    import type { Progress } from '@sourcegraph/shared/src/search/stream'

    export let progress: Progress

    const icons: Record<string, string> = {
        info: mdiInformationOutline,
        warning: mdiAlertCircle,
        error: mdiAlertCircle,
    }
    let searchAgainDisabled = true

    function updateButton(event: Event) {
        const element = event.target as HTMLInputElement
        searchAgainDisabled = Array.from(element.form?.querySelectorAll('[name="query"]') ?? []).every(
            input => !input.checked
        )
    }

    $: matchCount = progress.matchCount + (progress.skipped.length > 0 ? '+' : '')
    $: severity = progress.skipped.some(skipped => skipped.severity === 'warn' || skipped.severity === 'error')
        ? 'error'
        : 'info'
    $: hasSkippedItems = progress.skipped.length > 0
    $: sortedItems = sortBySeverity(progress.skipped)
    $: openItems = sortedItems.map((_, index) => index === 0)
    $: suggestedItems = sortedItems.filter(skipped => skipped.suggested)
    $: hasSuggestedItems = suggestedItems.length > 0
</script>

<Popover let:registerTrigger let:toggle placement="bottom-start">
    <button use:registerTrigger class="popover sm" on:click={() => toggle()}>
        <Icon svgPath={icons[severity]} inline />
        {matchCount} results in {(progress.durationMs / 1000).toFixed(2)}s
        <Icon svgPath={mdiChevronDown} inline />
    </button>
    <div slot="content" class="popover">
        <p>
            Found {limitHit(progress) ? 'more than ' : ''}
            {progress.matchCount}
            {pluralize('result', progress.matchCount)}
            {#if progress.repositoriesCount !== undefined}
                from {progress.repositoriesCount} {pluralize('repository', progress.repositoriesCount, 'repositories')}
            {/if}.
        </p>
        {#if hasSkippedItems}
            <h3>Some results skipped</h3>
            {#each sortedItems as item, index (item.reason)}
                {@const open = openItems[index]}
                <button class="item" on:click={() => (openItems[index] = !open)}>
                    <span>
                        <Icon svgPath={icons[item.severity]} inline />
                        {item.title}
                    </span>
                    {#if item.message}
                        <Icon svgPath={open ? mdiChevronDown : mdiChevronLeft} inline />
                    {/if}
                </button>
                {#if item.message && open}
                    <div class="message">
                        {@html renderMarkdown(item.message)}
                    </div>
                {/if}
            {/each}
        {/if}
        {#if hasSuggestedItems}
            <p>Search again:</p>
            <form on:submit|preventDefault>
                {#each suggestedItems as item (item.suggested.queryExpression)}
                    <label>
                        <input
                            type="checkbox"
                            name="query"
                            value={item.suggested.queryExpression}
                            on:click={updateButton}
                        />
                        {item.suggested.title} (
                        <SyntaxHighlightedQuery query={item.suggested.queryExpression} />
                        )
                    </label>
                {/each}
                <button class="mt-3" disabled={searchAgainDisabled}>
                    <Icon svgPath={mdiMagnify} />
                    <span>Search again</span>
                </button>
            </form>
        {/if}
    </div>
</Popover>

<style lang="scss">
    button {
        margin: 0;
        background-color: transparent;
        padding: 0.25rem 0.75rem;
        cursor: pointer;
        border: none;
        color: var(--body-color);
        text-align: left;

        &:disabled {
            cursor: not-allowed;
        }

        &.popover {
            border: 1px solid var(--border-color);
            border-radius: var(--border-radius);
        }

        &.sm {
            font-size: calc(min(0.75rem, 0.9166666667em));
            line-height: 1rem;
            letter-spacing: -0.0208333333em;
        }

        &.item {
            display: flex;
            justify-content: space-between;
            width: 100%;
        }
    }

    div.popover {
        width: 20rem;

        p,
        h3,
        form {
            margin: 1rem;
        }
    }

    div.message {
        border-left: 2px solid var(--primary);
        padding-left: 0.5rem;
        margin: 0 1rem 1rem 1rem;
    }

    label {
        display: block;
    }
</style>
