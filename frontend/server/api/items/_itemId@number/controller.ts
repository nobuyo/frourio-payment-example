import { getItem, getItems } from '$/service/items'
import { defineController } from './$relay'

export default defineController(() => ({
  get: async ({ params }) => {
    const item = await getItem(params.itemId)

    return item ? { status: 200, body: item } : { status: 404 }
  }
}))
