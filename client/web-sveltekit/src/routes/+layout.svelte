<script lang="ts">
    import Header from './Header.svelte'
    import './styles.scss'
    import type { LayoutData } from './$types'
    import { readable, writable } from 'svelte/store'
    import type { AuthenticatedUser } from '@sourcegraph/shared/src/auth'
    import type { Settings } from '@sourcegraph/shared/src/settings/settings'
    import { onMount, setContext } from 'svelte'
    import { KEY, type SourcegraphContext } from '$lib/stores'
    import { isErrorLike } from '@sourcegraph/common'
    import type { PlatformContext } from '@sourcegraph/shared/src/platform/context'
    import { observeSystemIsLightTheme } from '@sourcegraph/shared/src/theme'
    import { browser } from '$app/environment'
    import { readableObservable } from '$lib/utils'

    export let data: LayoutData

    const user = writable<AuthenticatedUser | null>(null)
    const settings = writable<Settings | null>(null)
    const platformContext = writable<PlatformContext | null>(null)
    const isLightTheme = browser ? readableObservable(observeSystemIsLightTheme(window).observable) : readable(true)

    setContext<SourcegraphContext>(KEY, {
        user,
        settings,
        platformContext,
        isLightTheme,
    })

    onMount(() => {
        // Settings can change over time. This ensures that the store is always
        // up-to-date.
        const settingsSubscription = data.platformContext?.settings.subscribe(settings => {
            $settings = isErrorLike(settings.final) ? null : settings.final
        })
        return () => settingsSubscription?.unsubscribe()
    })

    $: $user = data.user ?? null
    $: $settings = data.settings
    $: $platformContext = data.platformContext

    $: if (browser) {
        document.documentElement.classList.toggle('theme-light', $isLightTheme)
        document.documentElement.classList.toggle('theme-dark', !$isLightTheme)
    }

    $: console.log('settings', $settings)
</script>

<svelte:head>
    <title>Sourcegraph</title>
    <meta name="description" content="Code search" />
</svelte:head>

<div class="app">
    <Header authenticatedUser={$user} />

    <main>
        <slot />
    </main>
</div>

<style>
    .app {
        display: flex;
        flex-direction: column;
        height: 100vh;
        overflow: hidden;
    }

    main {
        flex: 1;
        display: flex;
        flex-direction: column;
        box-sizing: border-box;
        overflow: hidden;
    }
</style>
