import { get, post, upload } from '@/api/http'
import type { GetProp, UploadProps } from 'antd'
import { GLOBAL_STATUS } from './constants'

// /public/login
export type TLoginDetails = {
  refresh_token: string
}

export const login = params => post<TLoginDetails>('/public/login', params)

// /tokens/renew
export type TRenew = { access_token: string; access_token_expires_at: string }

export const renewToken = params => post<TRenew>('/tokens/renew', params)

export type TWebsiteItem = {
  id: string
  status: number
}

// /website/get
export const getWebsiteDetails = () => get<TWebsiteItem>('/website/get')

// /website/update
export const updateWebsiteDetails = params => post('/website/update', params)

//users/get
export type TDetailsItem = {
  address: string
  created_at: number
  description: string
  email: string
  first_name: string
  id: string
  last_name: string
  phone: string
  isdownloadable: number
  resume_pdf: string
  resume_docx: string
  updated_at: number
}

export const getInfo = () => get<TDetailsItem>('/users/get')

// /users/update
export const updateDetails = params => post<TDetailsItem>('/users/update', params)

// /links/getLinks
export type TLinksItem = {
  created_at: number
  id: string
  link: string
  updated_at: number
  type: string
  status: GLOBAL_STATUS
}
export const getLinks = () => get<TLinksItem[]>('/links/get')

// /links/addLinks
export const addLinks = params => post('/links/add', params)

// /links/deleteLinks
export const updateLinks = params => post('/links/update', params)

// /links/deleteLinks
export const deleteLinks = params => post('/links/delete', params)

// /links/toggleLinkStatus
export const toggleLinkStatus = params => post('/links/toggle', params)

// /services/get
export type TServiceItem = {
  id: string
  title: string
  description: string
  logo: string
  status: GLOBAL_STATUS
  created_at: number
  updated_at: number
}
export const getServices = () => get<TServiceItem[]>('/services/get')

// /services/update
export const updateServices = params => post('/services/update', params)

// /services/add
export const addServices = params => post('/services/add', params)

// /services/delete
export const deleteServices = params => post('/services/delete', params)

// /services/toggle
export const toggleServiceStatus = params => post('/services/toggle', params)

// /files/upload
export type TUploadImage = {
  url: string
  fileName: string
  size: number
}

export type FileType = Parameters<GetProp<UploadProps, 'beforeUpload'>>[0]

export const uploadImage = params => upload<{ url: string }>('/files/upload', params)

// /portfolios/add
export const addPortfolios = params => post('/portfolios/add', params)

// /portfolios/update
export const updatePortfolios = params => post('/portfolios/update', params)

// /portfolios/toggle
export const togglePortfolioStatus = params => post('/portfolios/toggle', params)

// /portfolios/delete
export const deletePortfolios = params => post('/portfolios/delete', params)

// /portfolios/get
export type TPortfolioItem = {
  id: string
  title: string
  tech: string[]
  link: string
  image: string
  status: GLOBAL_STATUS
  created_at: number
  updated_at: number
}
export const getPortfolios = () => get<TPortfolioItem[]>('/portfolios/get')

// /experiences/add
export const addExperiences = params => post('/experiences/add', params)

// /experiences/update
export const updateExperiences = params => post('/experiences/update', params)

// /experiences/delete
export const deleteExperiences = params => post('/experiences/delete', params)

// /experiences/get
export type TDescriptionItem = {
  id: string
  description: string
  experience_id: string
}

export type TExperienceItem = {
  id: string
  company: string
  title: string
  location: string
  started: string
  ended: string
  skills: { id: string; name: string; percentage: number }[]
  descriptions: string[]
  status: GLOBAL_STATUS
  created_at: number
  updated_at: number
}

export const getExperiences = () => get<TPortfolioItem[]>('/experiences/get')

// /experiences/toggle
export const toggleExperienceStatus = params => post('/experiences/toggle', params)

// /blogs/get
export type TBlogsItem = {
  title: string
  date: string
  description: string
  link: string
  image: string
  id: string
  status: GLOBAL_STATUS
  created_at: number
  updated_at: number
}
export const getBlogs = () => get<TBlogsItem[]>('/blogs/get')

// /blogs/add
export const addBlogs = params => post('/blogs/add', params)

// /blogs/update
export const updateBlogs = params => post('/blogs/update', params)

// /blogs/delete
export const deleteBlogs = params => post('/blogs/delete', params)

// /blogs/toggle
export const toggleBlogStatus = params => post('/blogs/toggle', params)

// /messages
export type TMessageItem = {
  id: string
  name: string
  email: string
  message: string
  created_at: number
  updated_at: number
}
export const getMessages = () => get<TMessageItem[]>('/messages/get')

// /messages/addMessages
export const addMessages = params => post('/public/sendMsg', params)

// /testimonials/get
export type TTestimonialsItem = {
  id: string
  author: string
  description: string
  image: string
  job: string
  status: GLOBAL_STATUS
  created_at: number
  updated_at: number
}
export const getTestimonials = () => get<TTestimonialsItem[]>('/testimonials/get')

// /testimonials/add
export const addTestimonials = params => post('/testimonials/add', params)

// /testimonials/update
export const updateTestimonials = params => post('/testimonials/update', params)

// /testimonials/delete
export const deleteTestimonials = params => post('/testimonials/delete', params)

// /testimonials/toggle
export const toggleTestimonialStatus = params => post('/testimonials/toggle', params)

// /educations/get
export type TEducationsItem = {
  id: string
  school: string
  course: string
  started: string
  ended: string
  description: string
  status: GLOBAL_STATUS
  skills: { id: string; name: string; percentage: number }[]
  created_at: number
  updated_at: number
}
export const getEducations = () => get<TEducationsItem[]>('/educations/get')

// /educations/add
export const addEducations = params => post('/educations/add', params)

// /educations/update
export const updateEducations = params => post('/educations/update', params)

// /educations/delete
export const deleteEducations = params => post('/educations/delete', params)

// /educations/toggle
export const toggleEducationStatus = params => post('/educations/toggle', params)

// /applications/get
export type TApplicationItem = {
  id: string
  name: string
  image: string
  status: number
  created_at: number
  updated_at: number
}
export const getApplication = params => get<TApplicationItem[]>('/applications/get', params)

// /applications/add
export const addApplication = params => post('/applications/add', params)

// /applications/update
export const updateApplication = params => post('/applications/update', params)

// /applications/delete
export const deleteApplication = params => post('/applications/delete', params)

// /applications/toggle
export const toggleAppStatus = params => post('/applications/toggle', params)

// /files/get
export type TFileItem = {
  id: string
  name: string
  file: string
  status: number
  created_at: string
}
export const getFiles = () => get<TFileItem[]>('/files/get')

export const deleteFile = params => post('/files/delete', params)

// /logs/get
export type TSessionItem = {
  id: string
  user_id: string
  email: string
  refresh_token: string
  is_revoked: boolean
  created_at: string
  expires_at: string
}
export const getLogs = () => get<TSessionItem[]>('/logs/get')
