import { useContext, useState } from 'react';
import { deleteDevice, Device, wakeDevice } from '../api';
import { DeviceContext } from '../DeviceContext';

export function DeviceCard({ id, name, mac }: Device) {
  const [confirmDelete, setConfirmDelete] = useState(false);
  const { removeDevice } = useContext(DeviceContext);

  const deletePc = () => {
    deleteDevice(id);
    removeDevice(id);
  };

  const wakePc = () => {
    wakeDevice(id);
  };

  return (
    <div className='bg-gray-300 my-3 p-3 rounded-md hover:bg-gray-200 hover:shadow-2xl transition-all'>
      <div className='flex justify-between'>
        <p className='font-bold text-2xl'>{name}</p>
        <div className='text-gray-200 text-lg'>
          <button
            className='bg-red-600 hover:bg-red-700 transition-color px-3 rounded-md mr-2'
            onClick={() => setConfirmDelete(!confirmDelete)}
          >
            Delete
          </button>
          <button
            className='bg-green-600 hover:bg-green-700 transition-color px-3 rounded-md'
            onClick={wakePc}
          >
            Wake
          </button>
        </div>
      </div>
      <div className='flex justify-between'>
        <p>{mac}</p>
        <p className='text-sm text-gray-500'>{id}</p>
      </div>

      {confirmDelete && (
        <>
          <p className='font-bold mt-5'>
            Are you sure you want to delete this entry?{' '}
            <span className='text-red-500'>This cannot be undone</span>
          </p>

          <button
            className='bg-red-600 hover:bg-red-700 transition-color text-gray-200 px-3 rounded-md mr-2'
            onClick={deletePc}
          >
            Yes, delete it
          </button>
          <button className='bg-green-600 hover:bg-green-700 transition-color text-gray-200 px-3 rounded-md mr-2'>
            No, keep it
          </button>
        </>
      )}
    </div>
  );
}
