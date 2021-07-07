import React, { ReactElement, useState } from 'react';
import { Select, SelectVariant } from '@patternfly/react-core';

export type SelectSingleProps = {
    id: string;
    value: string;
    handleSelect: (name: string, value: string) => void;
    isDisabled: boolean;
    children: ReactElement[];
    direction?: 'up' | 'down';
};

function SelectSingle({
    id,
    value,
    handleSelect,
    isDisabled,
    children,
    direction = 'down',
}: SelectSingleProps): ReactElement {
    const [isOpen, setIsOpen] = useState(false);

    function onSelect(_event, selection) {
        // The mouse event is not useful.
        setIsOpen(false);
        handleSelect(id, selection);
    }

    return (
        <Select
            variant={SelectVariant.single}
            id={id}
            isDisabled={isDisabled}
            isOpen={isOpen}
            onSelect={onSelect}
            onToggle={setIsOpen}
            selections={value}
            direction={direction}
        >
            {children}
        </Select>
    );
}

export default SelectSingle;