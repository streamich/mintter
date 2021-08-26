import {useCallback} from 'react'
import {useForm} from 'react-hook-form'
import {useMutation} from 'react-query'
import toast from 'react-hot-toast'

import {updateAccount} from '@mintter/client'
import {TextField} from '@mintter/ui'

import {
  OnboardingStep,
  OnboardingStepActions,
  OnboardingStepBody,
  OnboardingStepButton,
  OnboardingStepDescription,
  OnboardingStepTitle,
  ProfileInformationIcon,
} from './common'
import type {OnboardingStepPropsType} from './common'

type ProfileInformationDataType = {
  alias: string
  email: string
  bio: string
}

export function ProfileInformation({next}: OnboardingStepPropsType): JSX.Element {
  const updateProfile = useMutation(updateAccount)

  const {register, handleSubmit, errors, formState} = useForm<ProfileInformationDataType>({
    mode: 'onChange',
    defaultValues: {
      alias: '',
      email: '',
      bio: '',
    },
  })

  const onSubmit = useCallback(
    async (data: ProfileInformationDataType) => {
      await toast.promise(updateProfile.mutateAsync(data), {
        loading: 'Updating profile',
        success: 'Profile updated',
        error: 'Error updating profile',
      })
      console.log('edit complete!')
      next()
    },
    [next, updateProfile],
  )

  return (
    <OnboardingStep onSubmit={handleSubmit(onSubmit)}>
      <OnboardingStepTitle icon={<ProfileInformationIcon />}>Profile Information</OnboardingStepTitle>
      <OnboardingStepDescription>
        Link your personal data with your new Mintter account. You can fill this information later if you prefer.
      </OnboardingStepDescription>
      <OnboardingStepBody css={{display: 'flex', flexDirection: 'column', gap: '$6'}}>
        <TextField
          type="text"
          label="Alias"
          id="alias"
          name="alias"
          ref={register}
          data-testid="alias-input"
          placeholder="Readable alias or username. Doesn't have to be unique."
        />
        <TextField
          type="email"
          status={errors.email && 'danger'}
          label="Email"
          id="email"
          name="email"
          data-testid="email-input"
          ref={register({
            // pattern: {
            //   // eslint-disable-next-line no-control-regex
            //   value: /[^@ \t\r\n]+@[^@ \t\r\n]+\.[^@ \t\r\n]+/,
            //   message: 'Please type a valid email.',
            // },
          })}
          placeholder="Real email that could be publically shared"
          hint={errors.email?.message}
        />
        <TextField
          // TODO: fix types
          // @ts-ignore
          as="textarea"
          id="bio"
          name="bio"
          label="Bio"
          data-testid="bio-input"
          ref={register}
          rows={4}
          placeholder="A little bit about yourself..."
        />
      </OnboardingStepBody>
      <OnboardingStepActions>
        <OnboardingStepButton
          type="submit"
          data-testid="next-button"
          disabled={!formState.isValid || formState.isSubmitting}
        >
          Next
        </OnboardingStepButton>
      </OnboardingStepActions>
    </OnboardingStep>
  )
}