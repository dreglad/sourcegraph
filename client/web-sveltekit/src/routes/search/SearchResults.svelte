<script context="module" lang="ts">
    interface Cache {
        count: number
        expanded: Set<SearchMatch>
    }
    const cache = new Map<string, Cache>()

    const DEFAULT_INITIAL_ITEMS_TO_SHOW = 15
    const INCREMENTAL_ITEMS_TO_SHOW = 10
</script>

<script lang="ts">
    import type { AggregateStreamingSearchResults, SearchMatch } from '@sourcegraph/shared/src/search/stream'
    import SearchBox from '$lib/search/SearchBox.svelte'
    import type { Observable } from 'rxjs'
    import FileSearchResult from './FileSearchResult.svelte'
    import RepoSearchResult from './RepoSearchResult.svelte'
    import { beforeNavigate } from '$app/navigation'
    import { setContext, tick } from 'svelte'
    import { SearchPatternType } from '@sourcegraph/shared/src/schema'
    import { observeIntersection } from '$lib/intersection-observer'
    import Section from './SidebarSection.svelte'
    import { preserveScrollPosition } from '$lib/app'
    import LoadingSpinner from '$lib/LoadingSpinner.svelte'
    import { queryStateStore, submitSearch, type QueryOptions } from '$lib/search/state'
    import { settings } from '$lib/stores'
    import { searchTypes } from '$lib/search/sidebar'
    import SymbolSearchResult from './SymbolSearchResult.svelte'
    import StreamingProgress from './StreamingProgress.svelte'
    import type { SidebarFilter } from '$lib/search/utils'

    export let query: string
    export let stream: Observable<AggregateStreamingSearchResults | undefined>
    export let queryOptions: QueryOptions

    let patternType: SearchPatternType = SearchPatternType.literal
    let resultContainer: HTMLElement | null = null
    $: queryState = queryStateStore({ ...queryOptions, query }, $settings)

    $: loading = $stream && !$stream.progress.done
    $: results = $stream && $stream.results
    $: filters = $stream && $stream.filters
    $: langFilters =
        filters
            ?.filter(filter => filter.kind === 'lang')
            .map((filter): SidebarFilter => ({ ...filter, runImmediately: true })) ?? []

    // Logic for maintaining list state (scroll position, rendered items, open
    // items) for backwards navigation.
    $: cacheEntry = cache.get(query)
    $: count = cacheEntry?.count ?? DEFAULT_INITIAL_ITEMS_TO_SHOW
    $: resultsToShow = results ? results.slice(0, count) : null
    $: expandedSet = cacheEntry?.expanded || new Set<SearchMatch>()
    let scrollTop: number = 0
    preserveScrollPosition(
        position => (scrollTop = position ?? 0),
        () => resultContainer?.scrollTop
    )
    $: if (resultContainer) {
        resultContainer.scrollTop = scrollTop ?? 0
    }

    setContext('search-results', {
        isExpanded(match: SearchMatch): boolean {
            return expandedSet.has(match)
        },
        setExpanded(match: SearchMatch, expanded: boolean): void {
            if (expanded) {
                expandedSet.add(match)
            } else {
                expandedSet.delete(match)
            }
        },
    })
    beforeNavigate(() => {
        cache.set(query, { count, expanded: expandedSet })
    })

    $: progress = $stream?.progress

    function loadMore(event: { detail: boolean }) {
        if (event.detail) {
            count += INCREMENTAL_ITEMS_TO_SHOW
        }
    }

    async function updateQuery(event: MouseEvent) {
        const element = event.currentTarget as HTMLElement
        // TODO: Replace / update query; editor hints; etc
        queryState.setQuery(query => query + ' ' + element.dataset.value)
        if (element.dataset.run) {
            await tick()
            submitSearch($queryState)
        }
    }

    // FIXME: Not a great solution since it relies on implementation details of
    // the progress component
    async function onResubmitQuery(event: SubmitEvent) {
        const filters = Array.from(event.currentTarget?.querySelectorAll('[name="query"]') ?? [])
            .filter(input => input.checked)
            .map(input => input.value)
            .join(' ')
        queryState.setQuery(query => query + ' ' + filters)
        await tick()
        submitSearch($queryState)
    }
</script>

<section>
    <div class="search">
        <SearchBox {queryState} {patternType} selectedSearchContext={queryOptions.searchContext} />
    </div>

    <div class="results" bind:this={resultContainer}>
        <div class="scroll-container">
            {#if !$stream || loading}
                <div class="spinner">
                    <LoadingSpinner />
                </div>
            {:else if !loading && resultsToShow}
                <div class="main">
                    <aside class="stats mb-2">
                        {#if progress}
                            <StreamingProgress {progress} on:submit={onResubmitQuery} />
                        {/if}
                    </aside>
                    <ol>
                        {#each resultsToShow as result}
                            <li>
                                {#if result.type === 'content'}
                                    <FileSearchResult {result} />
                                {:else if result.type === 'repo'}
                                    <RepoSearchResult {result} />
                                {:else if result.type === 'symbol'}
                                    <SymbolSearchResult {result} />
                                {/if}
                            </li>
                        {/each}
                        <div use:observeIntersection on:intersecting={loadMore} />
                    </ol>
                </div>
                <aside class="sidebar">
                    <h4>Filters</h4>
                    <Section items={searchTypes} title="Search types" on:click={updateQuery} />
                    {#if langFilters.length > 1}
                        <Section items={langFilters} title="Languages" on:click={updateQuery} />
                    {/if}
                </aside>
            {/if}
        </div>
    </div>
</section>

<style lang="scss">
    .search {
        border-bottom: 1px solid var(--border-color);
        display: flex;
        align-self: stretch;
        padding: 0.5rem 1rem;
    }

    section {
        flex: 1;
        display: flex;
        align-items: center;
        flex-direction: column;
        overflow: hidden;

        :global(.search-box) {
            align-self: stretch;
        }
    }

    .results {
        flex: 1;
        align-self: stretch;
        overflow: auto;

        .scroll-container {
            padding: 1rem;
            display: flex;

            .spinner {
                flex: 1;
                display: flex;
                justify-content: center;
            }
        }
    }

    ol {
        padding: 0;
        margin: 0;
        list-style: none;
    }

    .main {
        flex: 1 1 auto;
        min-width: 0;
    }

    .sidebar {
        margin-left: 1rem;
        position: sticky;
        top: 1rem;
        align-self: flex-start;
        width: 15.5rem;
        flex-shrink: 0;
        background-color: var(--sidebar-bg);
        border: 1px solid var(--sidebar-border-color);
        padding: 0.75rem;
        border-radius: var(--border-radius);

        h4 {
            margin: -0.25rem -0.75rem 0.75rem;
            padding: 0 0.75rem 0.5rem;
            border-bottom: 1px solid var(--sidebar-border-color);
        }
    }
</style>
