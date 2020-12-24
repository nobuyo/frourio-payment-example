import Head from 'next/head'
import { useCallback, useState, FormEvent, ChangeEvent } from 'react'
import useAspidaSWR from '@aspida/swr'
import styles from '~/styles/ItemDetail.module.css'
import { apiClient } from '~/utils/apiClient'
import { Item, Purchase } from '$prisma/client'
import { useRouter } from 'next/router'

const ItemDetail = () => {
  const router = useRouter()
  const id = Number(router.query.id)
  const { data: items, error, revalidate } = useAspidaSWR(apiClient.items._itemId(id))

  if (error) return <div>failed to load</div>
  if (!items) return <div>loading...</div>

  return (
    <div>
      <Head>
        <title>EXAMPLE SHOP</title>
        <link rel="icon" href="/favicon.png" />
      </Head>

      <main>
        {id}
      </main>
    </div>
  )
}
export default ItemDetail
