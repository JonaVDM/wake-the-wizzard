export interface Device {
  id: string,
  name: string,
  mac: string,
}

export async function getDevices(): Promise<Device[]> {
  const req = await fetch('/api/pc');
  const body: Device[] = await req.json();
  return body;
}

export async function addDevice(name: string, mac: string) {
  fetch('/api/pc', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({name, mac})
  });
}

export async function deleteDevice(id: string) {
  fetch(`/api/pc/${id}`, { method: 'DELETE' });
}

export async function wakeDevice(id: string) {
  fetch(`/api/wake/${id}`);
}
