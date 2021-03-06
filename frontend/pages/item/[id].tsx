import Head from 'next/head'
import { useState, FormEvent } from 'react'
import useAspidaSWR from '@aspida/swr'
import styles from '~/styles/ItemDetail.module.css'
import { apiClient } from '~/utils/apiClient'
import { useRouter } from 'next/router'
import DatePicker from "react-datepicker"
import "react-datepicker/dist/react-datepicker.css"
import { normalizeKeys } from 'object-keys-normalizer';


const ItemDetail = () => {
  const router = useRouter()
  const id = Number(router.query.id)
  const { data: item, error, revalidate } = useAspidaSWR(apiClient.items._itemId(id))

  const [form, setForm] = useState({
    number: '',
    name: '',
    expYear: '',
    expMonth: '',
    cvc: '',
  });
  const [token, setToken] = useState('');
  const [expDate, setExpDate] = useState(new Date());
  const [purchased, setPurchased] = useState(false);

  if (error) return <div>failed to load</div>
  if (!item) return <div>loading...</div>

  const createToken = () => {
    const card = normalizeKeys(form, 'snake');
    window.Payjp.setPublicKey(process.env.payjpPublicKey);
    return new Promise((resolve, reject) => {
      window.Payjp.createToken(card, (status, res) => {
        if (status != 200) {
          window.alert(res.error.code || 'エラーが発生しました');
          reject(res.error.code);
        }
        setToken(res.id);
        resolve(res.id);
      });
    })
  }

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    await createToken();
    setTimeout(() => {}, 500);
    await apiClient.items._itemId(id).purchases.post({ body: { token } });
    const purchases = await apiClient.items._itemId(id).purchases.$get();
    setPurchased(!!purchases.length);

    window.alert('successfully purchased');

    router.push({
      pathname: '/',
    })
  }

  const handleChange = (input) => e => {
    setForm({ ...form, [input]: e.target.value });
  }

  const handleDateChange = (input: Date) => {
    setForm({ ...form, expMonth: input.getMonth().toString(), expYear: input.getFullYear().toString() });
    setExpDate(input);
  }

  return (
    <div>
      <Head>
        <title>EXAMPLE SHOP</title>
        <link rel="icon" href="/favicon.png" />
        <script src="https://js.pay.jp/"></script>
      </Head>

      <main>
        <div className={styles.container}>
          <a onClick={() => {
            router.push({
              pathname: '/',
            })
          }}>back</a>
          <hr/>
          <div className={styles.itemInfo}>
            <div className={styles.itemName}>{item?.name}</div>
            <div className={styles.itemDescription}>{item?.description}</div>
            <div className={styles.itemPrice}>{item?.amount}</div>
          </div>
          <form className={styles.cardForm} onSubmit={handleSubmit}>
            <div className={styles.formItem}>
              <label>Card Number</label>
              <input type="text" value={form.number} onChange={handleChange('number')} />
            </div>
            <div className={styles.formItem}>
              <label>Holder Name</label>
              <input type="text" value={form.name} onChange={handleChange('name')} />
            </div>
            <div className={styles.formItem}>
              <label>Exp</label>
              <DatePicker dateFormat="MM/yyyy" showMonthYearPicker selected={expDate} onChange={data => handleDateChange(data) } />
            </div>
            <div className={styles.formItem}>
              <label>CVC</label>
              <input type="password" value={form.cvc} onChange={handleChange('cvc')} />
            </div>
            <input type="submit" value="Submit" className={styles.submitButton} />
          </form>
        </div>
      </main>
    </div>
  )
}
export default ItemDetail
