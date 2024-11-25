import React, { createContext, useState, useEffect, useRef } from 'react';
export const ModalContext = createContext({isOpen: false, setIsOpen: undefined});

function ModalContextProvider({ children }) {
  const [isOpen, setIsOpen] = useState(false);
  const modalRef = useRef(null);

  useEffect(() => {
    if (isOpen) {
        return <div ref={modalRef}><div style={
          {
            background: 'blue',
            width: '100',
            height: '100',
            position: 'fixed',
            top: '0',
            left: '0',
            zIndex: '9999'
          }
        }>
        </div></div>
    }
  }, [isOpen])

  return (
    <ModalContext.Provider value={{ isOpen, setIsOpen }}>{children}</ModalContext.Provider>
  );
};

export default ModalContextProvider;