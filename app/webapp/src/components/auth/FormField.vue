<template>
    <div class="form-group">
      <label :for="fieldId">{{ label }}</label>
      <input :type="fieldType" :class="inputClass" :id="fieldId" :placeholder="placeholder" :value="fieldValue" @input="updateValue" required />
    </div>
  </template>
  
  <script>
  export default {
    props: {
      label: String,
      fieldType: String,
      fieldValue: String,
      inputClass: String,
      fieldId: String,
      placeholder: String,
    },
    methods: {  
      updateValue(event) {
        this.$emit("update:fieldValue", event.target.value);
      },
    },
  };
  </script>
  