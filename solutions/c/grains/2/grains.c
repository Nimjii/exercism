#include "grains.h"

uint64_t square(uint8_t index) {
   if (index < 1 || index > 64) {
      return 0;
   }

   uint64_t out = 1;

   return out << (index - 1);
}

uint64_t total(void) {
   uint64_t result = 0;

   for (int i = 1; i < 65; i++) {
      result += square(i);
   }

   return result;
}
