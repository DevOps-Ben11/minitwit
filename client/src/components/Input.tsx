import React from 'react'

type Props = {
  error?: string
} & React.InputHTMLAttributes<HTMLInputElement>

export const Input = React.forwardRef<HTMLInputElement, Props>(
  ({ error, type, ...props }, ref) => {
    return (
      <div className='input-component'>
        <input ref={ref} type={type || 'text'} {...props} />

        {error && <div className='input-component__error'>{error}</div>}
      </div>
    )
  }
)
