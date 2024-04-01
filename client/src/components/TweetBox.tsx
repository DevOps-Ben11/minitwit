import axios from 'axios'
import { SubmitHandler, useForm } from 'react-hook-form'
import { Input } from './Input'

type Props = {
  username: string
  handleCTA: () => void
}

type FormValues = {
  message: string
}

export const TweetBox = ({ username, handleCTA }: Props) => {
  const {
    register,
    reset,
    handleSubmit,
    formState: { errors },
  } = useForm<FormValues>()

  const onSubmit: SubmitHandler<FormValues> = async (data) => {
    try {
      await axios.post('/api/add_message', data)

      reset()
      handleCTA()
    } catch (error) {
      console.error(error)
    }
  }

  return (
    <div className='twitbox'>
      <h3>What's on your mind {username}?</h3>

      <form onSubmit={handleSubmit(onSubmit)} noValidate>
        <div className='twitbox-form'>
          <Input
            id='message'
            {...register('message', {
              required: 'Message is required',
            })}
            error={errors.message?.message}
          />

          <input type='submit' value='Share' />
        </div>
      </form>
    </div>
  )
}
