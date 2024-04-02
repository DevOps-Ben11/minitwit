import axios from 'axios'
import {
  UserTimelineResponse,
  PublicTimelineResponse,
  TimelineResponse,
} from './api.types'

export const getUserTimeline = async (username: string) =>
  axios.get<UserTimelineResponse>(`/api/timeline/${username}`)

export const getPublicTimeline = async () =>
  await axios.get<PublicTimelineResponse>('/api/public')

export const getTimeline = async () =>
  await axios.get<TimelineResponse>('/api/timeline')
