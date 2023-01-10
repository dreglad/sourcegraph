import * as React from 'react'

import classNames from 'classnames'

import { numberWithCommas } from '@sourcegraph/common'
import { LinkOrSpan } from '@sourcegraph/shared/src/components/LinkOrSpan'
import { Badge, Icon } from '@sourcegraph/wildcard'

import { Timestamp } from '../components/time/Timestamp'
import { GitRefFields } from '../graphql-operations'

import styles from './GitReference.module.scss'

export { queryGitReferences, REPOSITORY_GIT_REFS, gitReferenceFragments } from './loader'

export interface GitReferenceNodeProps {
    node: GitRefFields

    /** Link URL; if undefined, node.url is used. */
    url?: string

    /** Whether any ancestor element higher up in the tree is an `<a>` element. */
    ancestorIsLink?: boolean

    children?: React.ReactNode

    className?: string

    icon?: React.ComponentType<{ className?: string }>

    onClick?: React.MouseEventHandler<HTMLAnchorElement>
    nodeLinkClassName?: string

    ariaLabel?: string
}

export const GitReferenceNode: React.FunctionComponent<React.PropsWithChildren<GitReferenceNodeProps>> = ({
    node,
    url,
    ancestorIsLink,
    children,
    className,
    onClick,
    icon: ReferenceIcon,
    nodeLinkClassName,
    ariaLabel,
}) => {
    const mostRecentSig =
        node.target.commit &&
        (node.target.commit.committer && node.target.commit.committer.date > node.target.commit.author.date
            ? node.target.commit.committer
            : node.target.commit.author)
    const behindAhead = node.target.commit?.behindAhead
    url = url !== undefined ? url : node.url

    return (
        <li key={node.id} className={classNames('d-block list-group-item', styles.gitRefNode, className)}>
            <LinkOrSpan
                className={classNames(styles.gitRefNodeLink, nodeLinkClassName)}
                to={!ancestorIsLink ? url : undefined}
                onClick={onClick}
                data-testid="git-ref-node"
                aria-label={ariaLabel}
            >
                <span className="d-flex flex-wrap align-items-center">
                    {ReferenceIcon && <Icon className="mr-1" as={ReferenceIcon} aria-hidden={true} />}
                    {/*
                    a11y-ignore
                    Rule: "color-contrast" (Elements must have sufficient color contrast)
                    GitHub issue: https://github.com/sourcegraph/sourcegraph/issues/33343
                */}
                    <Badge className="a11y-ignore px-1 py-0 mr-2 text-break text-wrap text-justify" as="code">
                        {node.displayName}
                    </Badge>
                    {mostRecentSig && (
                        <small>
                            Updated <Timestamp date={mostRecentSig.date} />{' '}
                            {mostRecentSig.person && <>by {mostRecentSig.person.displayName}</>}
                        </small>
                    )}
                </span>
                {behindAhead && (
                    <small>
                        {numberWithCommas(behindAhead.behind)} behind, {numberWithCommas(behindAhead.ahead)} ahead
                    </small>
                )}
                {children}
            </LinkOrSpan>
        </li>
    )
}
