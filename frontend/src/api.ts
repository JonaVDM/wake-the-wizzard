const BASE = "/api/";

export async function getList(): Promise<string[]> {
  const request = await fetch(BASE + "list");
  const data = await request.json();
  return data;
}
