import { useAuth } from '@/lib/hooks/useAuth'
import axios from 'axios'
import { useState } from 'react'
import { useForm, SubmitHandler } from 'react-hook-form'
import { Input } from '@/components/Input'
import { PageWrapper } from '@/components/PageWrapper'
import { toast } from 'react-toastify'

type FormValues = {
  username: string
  email: string
  password: string
}

export const Login = () => {
  const { setUsername } = useAuth()
  const [error, setError] = useState<string | null>(null)
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormValues>()

  const onSubmit: SubmitHandler<FormValues> = async (data) => {
    setError(null)

    try {
      await axios.post('/api/login', data)
      setUsername(data.username)

      toast.success('You were logged in')
    } catch (error) {
      if (axios.isAxiosError(error) && error.response?.data?.error_msg) {
        setError(error.response.data.error_msg)
      }
    }
  }

  return (
    <PageWrapper>
      <h2>Sign in</h2>
      {error && (
        <div className='error'>
          <strong>Error:</strong> {error}
        </div>
      )}

      <form onSubmit={handleSubmit(onSubmit)} noValidate>
        <dl>
          <dt>
            <label htmlFor='username'>Username:</label>
          </dt>
          <dd>
            <Input
              id='username'
              {...register('username', {
                required: 'Invalid username',
              })}
              error={errors.username?.message}
            />
          </dd>
          <dt>
            <label htmlFor='password'>Password:</label>
          </dt>
          <dd>
            <Input
              id='password'
              type='password'
              {...register('password', {
                required: 'Invalid password',
              })}
              error={errors.password?.message}
            />
          </dd>
        </dl>

        <div className='actions'>
          <input type='submit' value='Sign Up' />
        </div>
      </form>
    </PageWrapper>
  )
}
