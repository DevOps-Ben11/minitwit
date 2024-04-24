import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useForm, SubmitHandler } from 'react-hook-form'
import { axios } from '@/services/api'
import { isAxiosError } from 'axios'
import { Input } from '@/components/Input'
import { PageWrapper } from '@/components/PageWrapper'
import { toast } from 'react-toastify'

type FormValues = {
  username: string
  email: string
  password: string
  passwordRepeat: string
}

export const Register = () => {
  const [error, setError] = useState<string | null>(null)
  const {
    register,
    handleSubmit,
    getValues,
    formState: { errors },
  } = useForm<FormValues>()

  const navigate = useNavigate()

  const onSubmit: SubmitHandler<FormValues> = async (data) => {
    setError(null)

    try {
      await axios.post('/api/register', data)

      toast.success('You were successfully registered and can login now')
      navigate('/login')
    } catch (error) {
      if (isAxiosError(error) && error.response?.data?.error_msg) {
        setError(error.response.data.error_msg)
      }
    }
  }

  return (
    <PageWrapper>
      <h2>Sign Up</h2>
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
                required: 'You have to enter a username',
              })}
              error={errors.username?.message}
            />
          </dd>

          <dt>
            <label htmlFor='email'>E-Mail:</label>
          </dt>
          <dd>
            <Input
              id='email'
              type='email'
              {...register('email', {
                required: 'You have to enter a valid email address',
                validate: (value) =>
                  value.includes('@') ||
                  'You have to enter a valid email address',
              })}
              error={errors.email?.message}
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
                required: 'You have to enter a password',
              })}
              error={errors.password?.message}
            />
          </dd>

          <dt>
            <label htmlFor='passwordRepeat'>
              Password <small>(repeat)</small>:
            </label>
          </dt>
          <dd>
            <Input
              id='passwordRepeat'
              type='password'
              {...register('passwordRepeat', {
                required: 'You have to enter a password',
                validate: (value) =>
                  value === getValues('password') ||
                  'The two passwords do not match',
              })}
              error={errors.passwordRepeat?.message}
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
