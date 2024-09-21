import { createSignal } from '@react-rxjs/utils';
import { merge } from 'rxjs';
import { v4 as uuid } from 'uuid';
import { DomainEvents } from '@domain/constants/DomainEvents';

export type DomainEvent = {
  id?: string;
  name: DomainEvents;
  data: {
    [key: string]: unknown;
  };
};

const [event$, setEvent] = createSignal<DomainEvent | undefined>();

export const eventMap$ = merge(event$);

const dispatch = (event: DomainEvent): void => {
  if (!event.id) {
    event.id = uuid();
  }
  setEvent(event);
  setEvent(undefined);
};

export default Object.freeze({
  dispatch
});
