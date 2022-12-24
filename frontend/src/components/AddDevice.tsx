import { useContext, useEffect, useState } from 'react';
import { addDevice } from '../api';
import { DeviceContext } from '../DeviceContext';

export function AddDevice() {
  const [name, setName] = useState<string>('');
  const [mac, setMac] = useState<string>('');
  const deviceContext = useContext(DeviceContext);

  const [nameError, setNameError] = useState<string>('');
  const [macError, setMacError] = useState<string>('');

  useEffect(() => {
    setName('');
    setMac('');
  }, []);

  const submit = async () => {
    setNameError('');
    setMacError('');

    if (validate()) {
      return;
    }

    await addDevice(name.trim(), mac.trim().toLowerCase());
  };

  const validate = (): boolean => {
    let hasError = false;

    if (name.trim() == '') {
      setNameError('Name cannot be empty');
      hasError = true;
    }

    if (mac.trim().length != 17) {
      // Should prob do a bit more validation in this, o well.
      setMacError('Invalid mac address');
      hasError = true;
    }

    return hasError;
  };

  return (
    <div className='bg-gray-300 my-3 p-3 rounded-md'>
      <label className='block italic'>Name</label>
      <input
        className='rounded-sm px-2 w-full'
        type='text'
        placeholder='Magic Rock'
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
      {nameError != '' && <p className='text-red-500'>{nameError}</p>}

      <label className='block italic mt-4'>Mac Address</label>
      <input
        className='rounded-sm px-2 w-full'
        type='text'
        placeholder='aa:bb:cc:dd:ee:ff'
        value={mac}
        onChange={(e) => setMac(e.target.value)}
      />
      {macError != '' && <p className='text-red-500'>{macError}</p>}

      <button
        className='rounded-md bg-blue-500 text-white mt-4 px-3 py-1'
        onClick={submit}
      >
        Add
      </button>
    </div>
  );
}
