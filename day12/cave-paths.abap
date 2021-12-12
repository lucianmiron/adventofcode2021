CLASS zml_cl_day12 DEFINITION
  PUBLIC
  FINAL
  CREATE PUBLIC .

  PUBLIC SECTION.
    INTERFACES if_oo_adt_classrun.
  PROTECTED SECTION.
  PRIVATE SECTION.
    TYPES: BEGIN OF mts_cave,
             name           TYPE string,
             adjiacentCaves TYPE REF TO data,
           END OF mts_cave,
           mtt_cave TYPE STANDARD TABLE OF mts_cave,
           BEGIN OF mts_double_cave,
             name TYPE string,
             occ  TYPE i,
           END OF mts_double_cave.

    METHODS: initialize,
      navigateToCave
        IMPORTING
          is_current_cave   TYPE mts_cave
          it_already_passed TYPE zml_day12_s_t
        RETURNING
          VALUE(rt_paths)   TYPE zml_day12_s_t.

    DATA: mc_easy_input   TYPE string,
          mc_better_input TYPE string,
          mc_input        TYPE string,
          mt_caves        TYPE mtt_cave.

ENDCLASS.



CLASS zml_cl_day12 IMPLEMENTATION.


  METHOD if_oo_adt_classrun~main.
    initialize( ).

    DATA: lt_caves TYPE mtt_cave.
    SPLIT mc_input AT cl_abap_char_utilities=>newline INTO TABLE DATA(lt_links).
    LOOP AT lt_links ASSIGNING FIELD-SYMBOL(<ls_link>).
      SPLIT <ls_link> AT '-' INTO TABLE DATA(lt_adj_caves).

      DATA: ls_cave TYPE mts_cave.

      READ TABLE lt_caves
      WITH KEY name = lt_adj_caves[ 1 ]
      ASSIGNING FIELD-SYMBOL(<ls_cave>).
      IF sy-subrc <> 0.
        ls_cave-name = lt_adj_caves[ 1 ].
        INSERT ls_cave INTO TABLE lt_caves.
        READ TABLE lt_caves WITH KEY name = ls_cave-name ASSIGNING <ls_cave>.
      ENDIF.

      FIELD-SYMBOLS: <lt_cave_table> TYPE STANDARD TABLE.
      CLEAR: ls_cave.
      IF <ls_cave>-adjiacentcaves IS INITIAL.
        CREATE DATA <ls_cave>-adjiacentcaves TYPE STANDARD TABLE OF mts_cave.
      ENDIF.
      ASSIGN <ls_cave>-adjiacentcaves->* TO <lt_cave_table>.

      READ TABLE lt_caves
        WITH KEY name = lt_adj_caves[ 2 ]
        ASSIGNING FIELD-SYMBOL(<ls_cave_adj>).
      IF sy-subrc = 0.
        INSERT <ls_cave_adj> INTO TABLE <lt_cave_table>.
      ELSE.
        APPEND INITIAL LINE TO lt_caves ASSIGNING FIELD-SYMBOL(<ls_cave_ref>).
        <ls_cave_ref>-name = lt_adj_caves[ 2 ].
        CREATE DATA <ls_cave_ref>-adjiacentcaves TYPE STANDARD TABLE OF mts_cave.
        INSERT <ls_cave_ref> INTO TABLE <lt_cave_table>.
      ENDIF.

      READ TABLE lt_caves
      WITH KEY name = lt_adj_caves[ 2 ]
      ASSIGNING FIELD-SYMBOL(<ls_cave_2>).
      ASSIGN <ls_cave_2>-adjiacentcaves->* TO <lt_cave_table>.
      INSERT <ls_cave> INTO TABLE <lt_cave_table>.
    ENDLOOP.

    mt_caves = lt_caves.

    DATA(paths) = navigatetocave(
                    EXPORTING
                      is_current_cave          = lt_caves[ 5 ]
                      it_already_passed = VALUE #( ) ).

    out->write( lines( paths ) ).
  ENDMETHOD.


  METHOD navigateToCave.

    IF is_current_cave-name = 'end'.
      rt_paths = VALUE #( ( is_current_cave-name ) ).
      RETURN.
    ENDIF.

    DATA lt_already_passed TYPE zml_day12_s_t.
    lt_already_passed = it_already_passed.

    DATA(lv_upper) = is_current_cave-name.
    TRANSLATE lv_upper TO UPPER CASE.
    IF lv_upper <> is_current_cave-name.
      INSERT is_current_cave-name INTO TABLE lt_already_passed.
    ENDIF.

    LOOP AT is_current_cave-adjiacentcaves->* ASSIGNING FIELD-SYMBOL(<ls_adj_cave>).
      DATA:ls_cave TYPE mts_cave.
      ls_cave = <ls_adj_cave>.
      IF ls_cave-name = 'start'.
        CONTINUE.
      ENDIF.
      DATA(lv_occurrences) = 0.
      LOOP AT lt_already_passed ASSIGNING FIELD-SYMBOL(<ls_already_passed>)
      WHERE table_line = ls_cave-name.
        lv_occurrences = lv_occurrences + 1.
      ENDLOOP.
      IF lv_occurrences = 2.
        CONTINUE.
      ELSEIF lv_occurrences = 1.
        " Allow it to be doubled only if any other cave is only once passed through
        DATA(lf_other_doubled) = abap_false.
        LOOP AT lt_already_passed ASSIGNING <ls_already_passed>.
          DATA(lv_search_occ) = 0.
          LOOP AT lt_already_passed TRANSPORTING NO FIELDS
          WHERE table_line = <ls_already_passed>.
            lv_search_occ = lv_search_occ + 1.
          ENDLOOP.

          IF lv_search_occ = 2.
            lf_other_doubled = abap_true.
            EXIT.
          ENDIF.
        ENDLOOP.
        IF lf_other_doubled = abap_true.
          CONTINUE.
        ENDIF.
      ENDIF.
      DATA(lt_paths) = navigatetocave(
        EXPORTING
          is_current_cave   = ls_cave
          it_already_passed = lt_already_passed ).
      LOOP AT lt_paths ASSIGNING FIELD-SYMBOL(<ls_path>).
        <ls_path> = |{ is_current_cave-name },{ <ls_path> }|.
      ENDLOOP.
      INSERT LINES OF lt_paths INTO TABLE rt_paths.
    ENDLOOP.
  ENDMETHOD.


  METHOD initialize.
    CONCATENATE:
'start-A' cl_abap_char_utilities=>newline
'start-b' cl_abap_char_utilities=>newline
'A-c    ' cl_abap_char_utilities=>newline
'A-b    ' cl_abap_char_utilities=>newline
'b-d    ' cl_abap_char_utilities=>newline
'A-end  ' cl_abap_char_utilities=>newline
'b-end  ' cl_abap_char_utilities=>newline
INTO mc_easy_input.

    CONCATENATE:
'dc-end'   cl_abap_char_utilities=>newline
'HN-start' cl_abap_char_utilities=>newline
'start-kj' cl_abap_char_utilities=>newline
'dc-start' cl_abap_char_utilities=>newline
'dc-HN'    cl_abap_char_utilities=>newline
'LN-dc'    cl_abap_char_utilities=>newline
'HN-end'   cl_abap_char_utilities=>newline
'kj-sa'    cl_abap_char_utilities=>newline
'kj-HN'    cl_abap_char_utilities=>newline
'kj-dc'    cl_abap_char_utilities=>newline
INTO mc_better_input.

    CONCATENATE:
'KF-sr'    cl_abap_char_utilities=>newline
'OO-vy'    cl_abap_char_utilities=>newline
'start-FP' cl_abap_char_utilities=>newline
'FP-end'   cl_abap_char_utilities=>newline
'vy-mi'    cl_abap_char_utilities=>newline
'vy-KF'    cl_abap_char_utilities=>newline
'vy-na'    cl_abap_char_utilities=>newline
'start-sr' cl_abap_char_utilities=>newline
'FP-lh'    cl_abap_char_utilities=>newline
'sr-FP'    cl_abap_char_utilities=>newline
'na-FP'    cl_abap_char_utilities=>newline
'end-KF'   cl_abap_char_utilities=>newline
'na-mi'    cl_abap_char_utilities=>newline
'lh-KF'    cl_abap_char_utilities=>newline
'end-lh'   cl_abap_char_utilities=>newline
'na-start' cl_abap_char_utilities=>newline
'wp-KF'    cl_abap_char_utilities=>newline
'mi-KF'    cl_abap_char_utilities=>newline
'vy-sr'    cl_abap_char_utilities=>newline
'vy-lh'    cl_abap_char_utilities=>newline
'sr-mi'    cl_abap_char_utilities=>newline
INTO mc_input.
  ENDMETHOD.
ENDCLASS.