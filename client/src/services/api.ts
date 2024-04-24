import axiosService from 'axios'
import {
  UserTimelineResponse,
  PublicTimelineResponse,
  TimelineResponse,
} from './api.types'
import { FormValues as RegisterBody } from '@/pages/Register'
import { FormValues as LoginBody } from '@/pages/Login'

export const axios = axiosService.create()

axios.interceptors.response.use(
  (response) => {
    // Any status code that lie within the range of 2xx cause this function to trigger
    // Do something with response data
    return response
  },
  function (error) {
    // Any status codes that falls outside the range of 2xx cause this function to trigger
    // Do something with response error

    // Handle logout
    if (error?.response?.status === 401) {
      window.location.href = '/logout'
    }

    return Promise.reject(error)
  }
)

export const registerUser = async (data: RegisterBody) =>
  await axios.post('/api/register', data)

export const loginUser = async (data: LoginBody) =>
  await axios.post('/api/login', data)

export const getUserTimeline = async (username: string) =>
  axios.get<UserTimelineResponse>(`/api/timeline/${username}`)

export const getPublicTimeline = async () =>
  await axios.get<PublicTimelineResponse>('/api/public')

export const getTimeline = async () =>
  await axios.get<TimelineResponse>('/api/timeline')

export const followUser = async (username: string) =>
  await axios.post(`/api/${username}/follow`)

export const unFollowUser = async (username: string) =>
  await axios.post(`/api/${username}/unfollow`)
