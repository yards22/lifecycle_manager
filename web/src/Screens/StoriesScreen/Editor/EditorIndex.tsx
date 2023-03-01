import { useCallback, useState } from 'react';
import { RichTextEditor } from '@mantine/rte';

const initialValue =
  '<p>Your initial <b>html value</b> or an empty string to init editor without value</p>';

export default function EditorIndex() {
  const handleImageUpload = useCallback(
    (file: File): Promise<string> =>
      new Promise((resolve, reject) => {
        const formData = new FormData();
        formData.append('image', file);

        fetch('https://api.imgbb.com/1/upload?key=api_key', {
          method: 'POST',
          body: formData,
        })
          .then((response) => response.json())
          .then((result) => resolve(result.data.url))
          .catch(() => reject(new Error('Upload failed')));
      }),
    []
  );

  const [value, onChange] = useState('');
  return (
    <RichTextEditor value={value} onChange={onChange} onImageUpload={handleImageUpload} id="rte" />
  );
}