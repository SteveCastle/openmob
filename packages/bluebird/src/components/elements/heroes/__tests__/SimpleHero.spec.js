import React from 'react';
import { shallow } from 'enzyme';
import SimpleHero from '../SimpleHero';

it('renders without crashing', () => {
  shallow(<SimpleHero title="Test title." />);
});
