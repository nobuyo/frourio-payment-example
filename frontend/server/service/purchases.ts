import { depend } from 'velona'
import { PrismaClient } from '@prisma/client'
import { Purchase, Prisma } from '$prisma/client'
import { getItem } from './items';

const prisma = new PrismaClient()

export const getPurchases = depend(
  { prisma: prisma as { purchase: { findMany(): Promise<Purchase[]> } } },
  async ({ prisma }, limit?: number) =>
    (await prisma.purchase.findMany()).slice(0, limit)
)

export const createPurchase = async (itemId: Purchase['itemId'], body: any = {}) => {
  const item = await getItem(itemId);
  if (!item) { return null; }

  prisma.purchase.create({ data: { amount: 2000, item: { connect: item } } })
}
