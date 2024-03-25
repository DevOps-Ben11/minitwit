import { useAuth } from '@/lib/hooks/useAuth'
import axios from 'axios'
import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { useForm, SubmitHandler } from 'react-hook-form'
import { Input } from '@/components/Input'

type FormValues = {
  username: string
  email: string
  password: string
}

const Login = () => {
  const { isAuthenticated, setUsername } = useAuth()
  const [error, setError] = useState<string | null>(null)
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormValues>()

  const navigate = useNavigate()

  const onSubmit: SubmitHandler<FormValues> = async (data) => {
    setError(null)

    try {
      await axios.post('/api/login', data)
      setUsername(data.username)

      navigate('/')
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
    </div>
  )
}

export default Login
