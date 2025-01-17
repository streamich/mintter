import {publishDraft, updateDraftV2} from '@app/client'
import {Editor} from 'slate'

export type EditorPageProps = {
  editor?: Editor
  shouldAutosave?: boolean
  publishDraft?: typeof publishDraft
  updateDraft?: typeof updateDraftV2
}

export type PublicationPageProps = {
  params?: {docId: string; version: string; blockId?: string}
}
