import { writable } from 'svelte/store';

export const devices = writable<Device[]>([]);
export const message = writable<string>();
export const error = writable<string>();

export async function loadDevices() {
  const req = await fetch('/api/pc');
  const body: Device[] = await req.json();
  devices.set(body);
}

export async function addDevice(name: string, mac: string) {
  const req = await fetch('/api/pc', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ name, mac })
  });

  const data = await req.json();
  if (req.status == 200) {
    devices.update(devices => [...devices, { name, mac, id: data.id }]);
    message.set('Device added!');
  } else {
    throw new Error(data.error);
  }
}

export async function deleteDevice(id: string) {
  devices.update(devices => devices.filter(device => device.id != id));
  await fetch(`/api/pc/${id}`, { method: 'DELETE' });
}

export async function wakeDevice(id: string) {
  try {
    const response = await fetch(`/api/wake/${id}`);
    const body = await response.text()
    switch (response.status) {
      case 200:
        message.set('Wake-On-Lan send!');
        break;

      case 404:
        error.set('PC not found');
        break;

      case 500:
        error.set(`Internal server error: ${body}`);
        break;

      default:
        error.set(`Error ${response.status}: ${body}`);
        break;
    }
  } catch (e) {
    console.log('somethig bad just happened: ', e);
  }
}
