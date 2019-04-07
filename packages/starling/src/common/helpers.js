const MILLISECONDS = 1000
const isObject = a => !!a && a.constructor === Object
const getValue = obj =>
  Object.entries(obj).reduce((acc, entry) => {
    if (entry[0] === 'seconds') {
      return new Date(entry[1] * MILLISECONDS).toString()
    }
    if (entry[0] === 'ID') {
      return entry[1]
    }
    return acc
  }, '')
export default obj => (isObject(obj) ? getValue(obj) : obj)
