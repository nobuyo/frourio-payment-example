import { depend } from 'velona'
import { PrismaClient } from '@prisma/client'
import { Purchase, Item } from '$prisma/client'
import { getItem } from './items';
import { PaymentManagerClient } from '../grpc/pay_grpc_pb';
import { ChargeRequest, ChargeResponse } from '../grpc/pay_pb';
import { rejects } from 'assert';
import { credentials } from 'grpc';

const prisma = new PrismaClient()
const client = new PaymentManagerClient("0.0.0.0:50051", credentials.createInsecure());

const requestPayment = async (item: Item, token: string): Promise<ChargeResponse> => {
  return new Promise((resolve, reject) => {
    const request = new ChargeRequest();
    request.setToken(token);
    request.setName(item?.name);
    request.setDesc(item?.description ?? '');
    request.setAmount(item?.amount);

    client.charge(request, (error, response) => {
      if (error) {
        reject(error);
      }
      resolve(response);
    });
  });
}

const capturePayment = async (chargeId: string): Promise<ChargeResponse> => {
  return new Promise((resolve, reject) => {
    const request = new ChargeRequest();
    request.setChargeId(chargeId);
    client.capture(request, (error, response) => {
      if (error) {
        reject(error);
      }
      resolve(response);
    });
  });
}

export const getPurchases = depend(
  { prisma: prisma as { purchase: { findMany(): Promise<Purchase[]> } } },
  async ({ prisma }, limit?: number) =>
    (await prisma.purchase.findMany()).slice(0, limit)
)

export const createPurchase = async (itemId: Purchase['itemId'], body: any = {}) => {
  const item = await getItem(itemId);
  if (!item) { return null; }

  try {
    const draft = await requestPayment(item, body.token);
    await capturePayment(draft.getId());
  } catch (error) {
    console.log(error);
  }

  return prisma.purchase.create({ data: { amount: item.amount, item: { connect: { id: item.id } } } });
}
