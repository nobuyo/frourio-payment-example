import { depend } from 'velona'
import { PrismaClient } from '@prisma/client'
import { Item, Prisma } from '$prisma/client'

const prisma = new PrismaClient()

export const getItems = depend(
  { prisma: prisma as { item: { findMany(): Promise<Item[]> } } },
  async ({ prisma }, limit?: number) =>
    (await prisma.item.findMany()).slice(0, limit)
)

export const getItem = (id: Item['id']) => prisma.item.findUnique({ where: { id } })
