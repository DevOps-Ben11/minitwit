type Props = {
  isLoading: boolean
  children: React.ReactNode
}

export const TimelineSkeleton = ({ isLoading, children }: Props) => {
  if (isLoading) {
    return (
      <ul className='timeline-skeleton'>
        {[...Array(10)].map((_, index) => (
          <MessageSkeleton height={Math.random()} key={index} />
        ))}
      </ul>
    )
  }

  return children
}

const MessageSkeleton = ({ height }: { height: number }) => (
  <li className='message-skeleton'>
    <div className='message-skeleton__avatar' />

    <div
      className='message-skeleton__content'
      style={{
        height: `${height * 20 + 10}px`,
      }}
    />
  </li>
)
