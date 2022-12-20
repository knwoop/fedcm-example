export const fetcher = async (
  resource: RequestInfo,
  init?: RequestInit,
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
): Promise<any> => {
  const res = await fetch(resource, init)

  if (!res.ok) {
    const errorRes = await res.json()
    throw new Error(errorRes.message ?? 'APIリクエスト中にエラーが発生しました')
  }

  return res.json()
}
