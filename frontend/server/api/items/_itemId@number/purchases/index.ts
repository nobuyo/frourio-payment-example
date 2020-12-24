import { Purchase } from '$prisma/client'

export type Methods = {
  get: {
    query?: {
      limit?: number
    }

    resBody: Purchase[]
  }
  post: {
    reqBody: {}
    resBody: Purchase
  }
}
