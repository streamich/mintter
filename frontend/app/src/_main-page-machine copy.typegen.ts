// This file was automatically generated. Edits will be overwritten

export interface Typegen0 {
  '@@xstate/typegen': true
  eventsCausingActions: {
    assignData: 'REPORT.DATA.SUCCESS'
  }
  internalEvents: {
    'xstate.init': {type: 'xstate.init'}
  }
  invokeSrcNameMap: {}
  missingImplementations: {
    actions: never
    services: never
    guards: never
    delays: never
  }
  eventsCausingServices: {}
  eventsCausingGuards: {}
  eventsCausingDelays: {}
  matchesStates: 'idle' | 'ready'
  tags: never
}
export interface Typegen1 {
  '@@xstate/typegen': true
  eventsCausingActions: {
    assignData: 'REPORT.DATA.SUCCESS'
  }
  internalEvents: {
    'xstate.init': {type: 'xstate.init'}
  }
  invokeSrcNameMap: {}
  missingImplementations: {
    actions: never
    services: never
    guards: never
    delays: never
  }
  eventsCausingServices: {}
  eventsCausingGuards: {}
  eventsCausingDelays: {}
  matchesStates: 'idle' | 'ready'
  tags: never
}
export interface Typegen2 {
  '@@xstate/typegen': true
  eventsCausingActions: {
    reconcileLibrary: 'RECONCILE'
    setEditorParams: 'goToEditor' | 'goToNew'
    setPublicationParams: 'goToNew' | ''
    pushHomeRoute: 'goHome'
    pushDraftRoute: 'goToEditor'
    pushPublicationRoute: ''
  }
  internalEvents: {
    '': {type: ''}
    'xstate.init': {type: 'xstate.init'}
    'done.invoke.router': {
      type: 'done.invoke.router'
      data: unknown
      __tip: 'See the XState TS docs to learn how to strongly type this.'
    }
    'error.platform.router': {type: 'error.platform.router'; data: unknown}
  }
  invokeSrcNameMap: {
    router: 'done.invoke.router'
    createNewDraft: 'done.invoke.main page.routes.createDraft:invocation[0]'
  }
  missingImplementations: {
    actions: 'reconcileLibrary'
    services: never
    guards: never
    delays: never
  }
  eventsCausingServices: {
    router: 'xstate.init'
    createNewDraft: 'goToNew' | 'createNewDraft'
  }
  eventsCausingGuards: {
    isPublication: 'goToNew'
    isDraft: 'goToNew'
    isPublicationValid: ''
  }
  eventsCausingDelays: {}
  matchesStates:
    | 'routes'
    | 'routes.idle'
    | 'routes.home'
    | 'routes.editor'
    | 'routes.publication'
    | 'routes.publication.validating'
    | 'routes.publication.valid'
    | 'routes.publication.error'
    | 'routes.settings'
    | 'routes.createDraft'
    | {
        routes?:
          | 'idle'
          | 'home'
          | 'editor'
          | 'publication'
          | 'settings'
          | 'createDraft'
          | {publication?: 'validating' | 'valid' | 'error'}
      }
  tags: 'topbar' | 'library' | 'sidepanel' | 'publication'
}