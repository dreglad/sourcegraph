import {RepoHeaderContributionsLifecycleProps} from "../RepoHeader";
import {HoverThresholdProps} from "../RepoContainer";
import {BatchChangesProps} from "../../batches";
import {ExternalLinkFields} from "../../graphql-operations";
import {BreadcrumbSetters} from "../../components/Breadcrumbs";
import {ResolvedRevision} from "../backend";
import {CodeIntelligenceProps} from "../../codeintel";
import {StreamingSearchResultsListProps} from "@sourcegraph/search-ui/out/src";
import {ErrorLike} from "@storybook/client-api";
import {AuthenticatedUser} from "../../auth";
import {CodeInsightsProps} from "../../insights/types";
import {SearchStreamingProps} from "../../search";
import {SettingsCascadeProps} from "@sourcegraph/shared/out/src/settings/settings";
import {ActionItemsBarProps} from "../../extensions/components/ActionItemsBar";
import {PlatformContextProps} from "@sourcegraph/shared/out/src/platform/context";
import {TelemetryProps} from "@sourcegraph/shared/out/src/telemetry/telemetryService";
import {ExtensionsControllerProps} from "@sourcegraph/shared/out/src/extensions/controller";
import {RepoSettingsAreaRoute} from "./RepoSettingsArea";
import {RepoSettingsSideBarGroup} from "./RepoSettingsSidebar";
import {SearchContextProps} from "@sourcegraph/shared/out/src/search";
import {ThemeProps} from "@sourcegraph/shared/out/src/theme";
import {RouteDescriptor} from "../../util/contributions";

/** A sub-route of {@link RepoContainer}. */
export interface RepoSettingsContainerRoute extends RouteDescriptor<RepoSettingsContainerContext> {}

export interface RepoSettingsContainerContext
    extends RepoHeaderContributionsLifecycleProps,
        SettingsCascadeProps,
        ExtensionsControllerProps,
        PlatformContextProps,
        ThemeProps,
        HoverThresholdProps,
        TelemetryProps,
        Pick<SearchContextProps, 'selectedSearchContextSpec' | 'searchContextsEnabled'>,
        BreadcrumbSetters,
        ActionItemsBarProps,
        SearchStreamingProps,
        Pick<StreamingSearchResultsListProps, 'fetchHighlightedFileLineRanges'>,
        CodeIntelligenceProps,
        BatchChangesProps,
        CodeInsightsProps {
    // repo: RepositoryFields | undefined
    // repo: RepositoryFields
    repoName: string
    resolvedRevisionOrError: ResolvedRevision | ErrorLike | undefined
    authenticatedUser: AuthenticatedUser | null
    repoSettingsAreaRoutes: readonly RepoSettingsAreaRoute[]
    repoSettingsSidebarGroups: readonly RepoSettingsSideBarGroup[]

    /** The URL route match for {@link RepoContainer}. */
    routePrefix: string

    onDidUpdateExternalLinks: (externalLinks: ExternalLinkFields[] | undefined) => void

    globbing: boolean

    isMacPlatform: boolean

    isSourcegraphDotCom: boolean
}
