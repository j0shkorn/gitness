import React, { useMemo } from 'react'
import { Container, Color, TableV2 as Table, Text, Avatar, Layout } from '@harness/uicore'
import type { CellProps, Column } from 'react-table'
import { orderBy } from 'lodash-es'
import { useStrings } from 'framework/strings'
import { useAppContext } from 'AppContext'
import type { RepoCommit } from 'services/code'
import { CommitActions } from 'components/CommitActions/CommitActions'
import { NoResultCard } from 'components/NoResultCard/NoResultCard'
import { ThreadSection } from 'components/ThreadSection/ThreadSection'
import { formatDate } from 'utils/Utils'
import { CodeIcon, GitInfoProps } from 'utils/GitUtils'
import css from './CommitsView.module.scss'

interface CommitsViewProps extends Pick<GitInfoProps, 'repoMetadata'> {
  commits: RepoCommit[]
  emptyTitle: string
  emptyMessage: string
}

export function CommitsView({ repoMetadata, commits, emptyTitle, emptyMessage }: CommitsViewProps) {
  const { getString } = useStrings()
  const { routes } = useAppContext()
  const columns: Column<RepoCommit>[] = useMemo(
    () => [
      {
        id: 'author',
        width: '20%',
        Cell: ({ row }: CellProps<RepoCommit>) => {
          return (
            <Layout.Horizontal spacing="small" flex={{ alignItems: 'center' }} style={{ display: 'inline-flex' }}>
              <Avatar hoverCard={false} size="small" name={row.original.author?.identity?.name || ''} />
              <Text className={css.rowText} color={Color.BLACK}>
                {row.original.author?.identity?.name}
              </Text>
            </Layout.Horizontal>
          )
        }
      },
      {
        id: 'commit',
        width: 'calc(80% - 100px)',
        Cell: ({ row }: CellProps<RepoCommit>) => {
          return (
            <Text color={Color.BLACK} lineClamp={1} className={css.rowText}>
              {row.original.message}
            </Text>
          )
        }
      },
      {
        id: 'sha',
        width: '100px',
        Cell: ({ row }: CellProps<RepoCommit>) => {
          return (
            <CommitActions
              sha={row.original.sha as string}
              href={routes.toCODECommits({
                repoPath: repoMetadata.path as string,
                commitRef: row.original.sha as string
              })}
              enableCopy
            />
          )
        }
      }
    ],
    [repoMetadata.path, routes]
  )
  const commitsGroupedByDate: Record<string, RepoCommit[]> = useMemo(
    () =>
      commits?.reduce((group, commit) => {
        const date = formatDate(commit.author?.when as string)
        group[date] = (group[date] || []).concat(commit)
        return group
      }, {} as Record<string, RepoCommit[]>) || {},
    [commits]
  )

  return (
    <Container className={css.container}>
      {!!commits.length &&
        Object.entries(commitsGroupedByDate).map(([date, commitsByDate]) => {
          return (
            <ThreadSection
              key={date}
              title={
                <Text icon={CodeIcon.Commit} iconProps={{ size: 20 }} color={Color.GREY_500} className={css.label}>
                  {getString('commitsOn', { date })}
                </Text>
              }>
              <Table<RepoCommit>
                className={css.table}
                hideHeaders
                columns={columns}
                data={orderBy(commitsByDate || [], ['author.when'], ['desc'])}
                getRowClassName={() => css.row}
              />
            </ThreadSection>
          )
        })}
      <NoResultCard
        showWhen={() => commits?.length === 0}
        forSearch={true}
        title={emptyTitle}
        emptySearchMessage={emptyMessage}
      />
    </Container>
  )
}
