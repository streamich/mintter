import {ErrorBoundary} from 'react-error-boundary'
import {Link} from 'wouter'
import {useFiles} from '../../main-page-context'
import {PublicationRef} from '../../main-page-machine'
import {Section} from './section'
import {SectionError} from './section-error'
import {SectionItem} from './section-item'

export function FilesSection() {
  const files = useFiles()
  // let files = []
  return (
    <Section title="Files" open={true}>
      {!!files.length ? (
        <ErrorBoundary
          FallbackComponent={SectionError}
          onReset={() => {
            window.location.reload()
          }}
        >
          {files.map((publication: PublicationRef) => {
            let {ref, document} = publication
            console.log('document item: ', document)

            return (
              <Link href={`/p/${document?.id}`}>
                <SectionItem key={document?.id} href={`/p/${document?.id}`} document={document} />
              </Link>
            )
          })}
        </ErrorBoundary>
      ) : null}
    </Section>
  )
}