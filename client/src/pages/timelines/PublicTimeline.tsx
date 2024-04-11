import { TimelineSkeleton } from '@/components/TimelineSkeleton'
import { MessageList } from '@/components/MessageList'
import { PageWrapper } from '@/components/PageWrapper'
import { getPublicTimeline } from '@/services/api'
import { Message } from '@/services/api.types'
import { useEffect, useState } from 'react'

export const PublicTimeline = () => {
  const [messages, setMessages] = useState<Message[]>([])
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    const fetchMessages = async () => {
      try {
        setIsLoading(true)
        const response = await getPublicTimeline()
        setMessages(response.data.Messages)
      } catch (error) {
        console.error(error)
      } finally {
        setIsLoading(false)
      }
    }

    fetchMessages()
  }, [])

  return (
    <PageWrapper>
      <h2>Public Timeline</h2>

      <TimelineSkeleton isLoading={isLoading}>
        <MessageList messages={messages} />
      </TimelineSkeleton>
    </PageWrapper>
  )
}
