// eslint-disable-next-line
import {PublicationRef} from '@app/main-machine'
import {createTippingMachine} from '@app/tipping-machine'
import {Box} from '@components/box'
import {Button} from '@components/button'
import {Icon} from '@components/icon'
import {Text} from '@components/text'
import {TextField} from '@components/text-field'
import * as PopoverPrimitive from '@radix-ui/react-popover'
import {useActor, useInterpret} from '@xstate/react'
import QRCode from 'react-qr-code'
import {StateFrom} from 'xstate'

export function TippingModal({fileRef}: {fileRef: PublicationRef}) {
  // if (!visible) return null

  const [fileState] = useActor(fileRef)
  const service = useInterpret(
    createTippingMachine({
      accountID: fileState.context.author?.id,
      publicationID: fileState.context.documentId,
    }),
  )
  const [state, send] = useActor(service)

  if (typeof fileRef == 'undefined') {
    return null
  }

  return (
    <PopoverPrimitive.Root
      open={state.matches('open')}
      onOpenChange={(newVal) => {
        if (newVal) {
          send('OPEN')
        } else {
          send('CLOSE')
        }
      }}
    >
      <PopoverPrimitive.Trigger asChild>
        <Button
          color="success"
          size="1"
          variant="ghost"
          disabled={state.hasTag('pending')}
          data-testid="submit-edit"
          onClick={() => {
            send('OPEN')
          }}
          css={{
            display: 'flex',
            fontSize: '$2',
            flexDirection: 'row',
            alignItems: 'center',
            justifyContent: 'start',
            fontFamily: '$base',
            gap: '$4',
            paddingVertical: '$3',
            paddingHorizontal: '$4',
            borderRadius: '$2',
            '&[data-disabled]': {
              cursor: 'default',
            },
            cursor: 'pointer',
            '&:focus': {
              outline: 'none',
              backgroundColor: '$base-component-bg-normal',
              cursor: 'pointer',
            },
            '&:disabled': {
              opacity: 0.5,
            },
          }}
        >
          <Icon name="Star" />
          Tip {fileState.context.author?.profile?.alias}
        </Button>
      </PopoverPrimitive.Trigger>
      <PopoverPrimitive.Content>
        {state.matches('open.setAmount') && (
          <SetAmount state={state} send={send} />
        )}
        {state.matches('open.requestInvoice') ||
          (state.matches('open.paying') && (
            <Box
              css={{
                padding: '$5',
                width: '300px',
                backgroundColor: '$base-component-bg-normal',
                display: 'flex',
                flexDirection: 'column',
                gap: '$4',
                boxShadow: '$3',
              }}
            >
              <Text>...</Text>
            </Box>
          ))}
        {state.matches('open.errored') && (
          <Box
            css={{
              padding: '$5',
              width: '300px',
              backgroundColor: '$base-component-bg-normal',
              display: 'flex',
              flexDirection: 'column',
              gap: '$4',
              boxShadow: '$3',
            }}
          >
            <Text>Error:</Text>
            <Text size="1" color="danger">
              {JSON.stringify(state.context.errorMessage)}
            </Text>
            <Button
              size="1"
              type="submit"
              css={{width: '$full'}}
              onClick={() => send('RETRY')}
            >
              Retry
            </Button>
          </Box>
        )}
        {state.matches('open.readyToPay') && (
          <Box
            css={{
              padding: '$5',
              width: '300px',
              backgroundColor: '$base-component-bg-normal',
              display: 'flex',
              flexDirection: 'column',
              gap: '$4',
              boxShadow: '$3',
              svg: {
                width: '100%',
              },
            }}
          >
            <QRCode
              title="demo demo"
              value={state.context.invoice}
              size={300 - 32}
            />
            <Box>
              <Text size="1" fontWeight="bold">
                Invoice:
              </Text>
              <Text
                size="1"
                css={{wordBreak: 'break-all', wordWrap: 'break-word'}}
              >
                {state.context.invoice}
              </Text>
            </Box>
            <Button
              size="1"
              css={{width: '$full'}}
              onClick={() => send('TIPPING.PAY.INVOICE')}
            >
              Pay Directly
            </Button>
          </Box>
        )}
        {state.matches('open.success') && (
          <Box
            css={{
              padding: '$5',
              width: '300px',
              backgroundColor: '$base-component-bg-normal',
              display: 'flex',
              flexDirection: 'column',
              gap: '$4',
              boxShadow: '$3',
            }}
          >
            <Icon name="Star" />
            <Text>Payment Success!</Text>
          </Box>
        )}
      </PopoverPrimitive.Content>
    </PopoverPrimitive.Root>
  )
}

export function SetAmount({
  send,
  state,
}: {
  state: StateFrom<ReturnType<typeof createTippingMachine>>
  send
}) {
  return (
    <Box
      css={{
        padding: '$5',
        width: '300px',
        backgroundColor: '$base-component-bg-normal',
        display: 'flex',
        flexDirection: 'column',
        gap: '$4',
        boxShadow: '$3',
      }}
    >
      <Text size="4">Tip this Author</Text>
      {
        <Box css={{display: 'flex', flexDirection: 'column', gap: '$3'}}>
          <TextField
            type="number"
            id="amount"
            name="amount"
            label="Invoice Amount"
            size={1}
            value={state.context.amount}
            onChange={(e) =>
              send({
                type: 'TIPPING.UPDATE.AMOUNT',
                amount: Number(e.target.value),
              })
            }
          />
          <Box
            css={{
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'space-between',
            }}
          >
            <Button
              size="1"
              type="submit"
              disabled={state.hasTag('pending')}
              css={{width: '$full'}}
              onClick={() => send('TIPPING.REQUEST.INVOICE')}
            >
              Request Invoice
            </Button>
          </Box>
        </Box>
      }
    </Box>
  )
}
