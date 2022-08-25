<template>
  <div class="t-form">
    <t-form ref="form" :data="formData" label-width="calc(2em + 24px)" :layout="layout" :rules="rules" scroll-to-first-error="smooth">
      <template v-for="(item, index) in formItems" :key="index">
        <t-form-item :label="item.label" v-if="item.type === 'input'">
          <t-input v-model="formData[`${item.field}`]" :placeholder="item.placeholder"></t-input>
        </t-form-item>
        <t-form-item :label="item.label" v-else-if="item.type === 'select'">
          <t-select v-model="formData[`${item.field}`]" :placeholder="item.placeholder">
            <t-option v-for="(option, index) in item.options" :key="index" :value="option.value" :label="option.label"></t-option>
          </t-select>
        </t-form-item>
        <t-form-item :label="item.label" v-else-if="item.type === 'checkbox'">
          <t-checkbox-group v-model="formData[`${item.field}`]" :options="item.options"></t-checkbox-group>
        </t-form-item>
      </template>
      <t-form-item>
        <slot name="btn">

        </slot>
      </t-form-item>
    </t-form>
  </div>
</template>

<script setup lang="ts">
import { SubmitContext } from 'tdesign-vue-next'
import { defineProps, PropType, ref, watch } from 'vue'
import { IFormItem } from '../types'
const props = defineProps({
  modelValue: {
    type: Object,
    required: true,
  },
  formItems: {
    type: Array as PropType<IFormItem[]>,
    default: () => [],
  },
  rules: {
    type: Object,
    default: () => ({}),
  },
  itemStyle: {
    type: Object,
    default: () => ({}),
  },
  layout:{
    type: String,
    default: 'inline',
  }
})
const form = ref(null)

const emit = defineEmits(['update:modelValue', 'reset', 'submit'])
const formData = ref({ ...props.modelValue })
watch(
  () => props.modelValue,
  (val) => {
    formData.value = { ...val }
  }
)
watch(
  formData,
  (val) => {
    emit('update:modelValue', val)
  },
  { deep: true }
)
</script>

<style scoped>
.t-form {
  @apply p-4 bg-white;
}
</style>
