import {useFiles} from '@app/main-page-context'
import {PublicationRef} from '@app/main-page-machine'
import {LibraryItem} from '@components/library/library-item'
import {ErrorBoundary} from 'react-error-boundary'
import {Section} from './section'
import {SectionError} from './section-error'

export function FilesSection() {
  const files = useFiles()

  return (
    <Section title="Files" open={true}>
      {files.length ? (
        <ErrorBoundary
          FallbackComponent={SectionError}
          onReset={() => {
            window.location.reload()
          }}
        >
          {files.map((publication: PublicationRef) => {
            let {document, version} = publication
            return <LibraryItem key={document?.id} href={`/p/${document?.id}/${version}`} publication={publication} />
          })}
        </ErrorBoundary>
      ) : null}
    </Section>
  )
}