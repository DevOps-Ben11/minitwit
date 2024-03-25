import { useEffect, useState } from 'react'
import '../style.css'
import { useNavigate } from 'react-router-dom'
import { useForm, SubmitHandler } from 'react-hook-form'
import { useAuth } from '@/lib/hooks/useAuth'
import axios from 'axios'
import { Input } from '@/components/Input'

type FormValues = {
  username: string
  email: string
  password: string
  passwordRepeat: string
}

const Register = () => {
  const { isAuthenticated } = useAuth()
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
      navigate('/login')
    } catch (error) {
      if (axios.isAxiosError(error) && error.response?.data?.error_msg) {
        setError(error.response.data.error_msg)
      }
    }
  }

  useEffect(() => {
    if (isAuthenticated) {
      navigate('/')
    }
  }, [isAuthenticated, navigate])
  if (isAuthenticated) return null

  return (
    <div>
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
    </div>
  )
}

export default Register
