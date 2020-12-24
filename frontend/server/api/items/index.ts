import { Item } from '$prisma/client'

export type Methods = {
  get: {
    query?: {
      limit?: number
    }

    resBody: Item[]
  }
  // post: {
  //   reqBody: {}
  //   resBody: Item
  // }
}
