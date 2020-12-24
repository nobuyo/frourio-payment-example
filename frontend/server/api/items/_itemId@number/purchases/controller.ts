import { createPurchase, getPurchases } from '$/service/purchases'
import { Purchase } from '@prisma/client'
import { defineController } from './$relay'

export default defineController(() => ({
  get: async ({ query }) => {
    return { status: 200, body: await getPurchases(query?.limit) }
  },
  post: async ({ params, body }) => {
    const purchase = await createPurchase(params.itemId, body)
    if (!purchase) {
      return { status: 400 }
    }

    return { status: 201, body: purchase }
  }
}))
