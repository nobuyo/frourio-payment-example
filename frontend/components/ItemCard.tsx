import styles from '~/styles/ItemCard.module.css'

const ItemCard = (props) => {
  return (
    <div className={styles.itemCard} onClick={ props.onClick }>
      <div className={styles.itemName}>{ props.item?.name }</div>
      <div className={styles.itemDesc}>{ props.item?.description }</div>
      <div className={styles.itemAmount}>JPY { props.item?.amount }</div>
    </div>
  )
}

export default ItemCard
