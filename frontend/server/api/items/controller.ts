import { getItems } from '$/service/items'
import { defineController } from './$relay'

export default defineController(() => ({
  get: async ({ query }) => ({
    status: 200,
    body: (await getItems()).slice(0, query?.limit)
  })
}))
