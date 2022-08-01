export const fetcher = async (
  resouce: RequestInfo,
  init?: RequestInit,
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
): Promise<any> => {
  const res = await fetch(resouce, init)
  if (!res.ok) {
    const errorRes = await res.json()
    const error = new Error(errorRes.message ?? 'The ERROR happend in API call')
    throw error
  }
  return res.json()
}
