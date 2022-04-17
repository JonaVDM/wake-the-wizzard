const BASE = "/api/";

export async function getList(): Promise<string[]> {
  const request = await fetch(BASE + "list");
  const data = await request.json();
  return data;
}

export async function wakePc(mac: string) {
  const request = await fetch(BASE + `wake/${mac}`);
  const data = await request.text();
  return data === "ok";
}
