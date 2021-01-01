import Head from 'next/head'
import useAspidaSWR from '@aspida/swr'
import styles from '~/styles/Home.module.css'
import { apiClient } from '~/utils/apiClient'
import ItemCard from '~/components/ItemCard'
import { useRouter } from 'next/router'

const Home = () => {
  const { data: items, error, revalidate } = useAspidaSWR(apiClient.items)
  const router = useRouter()

  if (error) return <div>failed to load</div>
  if (!items) return <div>loading...</div>

  return (
    <div className={styles.container}>
      <Head>
        <title>EXAMPLE SHOP</title>
        <link rel="icon" href="/favicon.png" />
      </Head>

      <main className={styles.main}>
        <p className={styles.description}>gRPC + Frourio + Prisma app</p>

        <div>
          <ul className={styles.tasks}>
            {items.map((item) => (
              <ItemCard item={item} key={item.id} onClick={() => {
                router.push({
                  pathname: '/item/[id]',
                  query: { id: item.id },
                })
              }}></ItemCard>
            ))}
          </ul>
        </div>
      </main>
    </div>
  )
}

export default Home
