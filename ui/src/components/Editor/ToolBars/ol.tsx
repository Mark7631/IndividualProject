import { memo } from 'react';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

let context: IEditorContext;
const OL = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });
  const item = {
    label: 'orderedList',
    keyMap: ['Ctrl-O'],
    tip: `${t('ordered_list.text')} (Ctrl+O)`,
  };

  const handleClick = (ctx) => {
    context = ctx;
    const { editor, replaceLines } = context;

    replaceLines((line, i) => {
      const FIND_OL_RX = /^(\s{0,})(\d+)\.\s/;

      if (line.match(FIND_OL_RX)) {
        line = line.replace(FIND_OL_RX, '');
      } else {
        line = `${i + 1}. ${line}`;
      }
      return line;
    });
    editor.focus();
  };

  return <ToolItem {...item} onClick={handleClick} />;
};

export default memo(OL);
