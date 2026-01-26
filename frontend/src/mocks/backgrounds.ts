export const __mock__backgroundImages = [
  "https://cdn.hytale.com/5e7b9ecb50cbcd001176c5c1_11___z2_camels.png",
  "https://cdn.hytale.com/5e7b9f7e50cbcd001176c5d3_14___zone_3_sunshaft_and_bloom.jpg",
  "https://cdn.hytale.com/5e7b9f8a50cbcd001176c5d7_16___library_at_dusk.jpg",
  "https://cdn.hytale.com/5e7b9f9150cbcd001176c5db_18___cave_sunshaft.jpg",
  "https://cdn.hytale.com/5e7ba02d50cbcd001176c5ff_30___outlander_settlement.jpg",
  "https://cdn.hytale.com/5e7ba11250cbcd001176c64f_58___woodpecker.jpg",
  "https://cdn.hytale.com/5e7ba2bd50cbcd001176c6ad_87___ram.jpg",
  "https://cdn.hytale.com/5e7ba3f03c9a2a0010679360_90___z3_cave.jpg"
]

export const getRandomBgImage = (): string => {
  return __mock__backgroundImages[Math.floor(Math.random() * __mock__backgroundImages.length)]
}